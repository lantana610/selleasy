from flask import Flask, request, jsonify
from email.mime.text import MIMEText
import os
import smtplib

app = Flask(__name__)

SMTP_HOST = "smtp.gmail.com"
SMTP_PORT = 587
SMTP_USER = os.environ.get("SMTP_USER", "")
SMTP_PASSWORD = os.environ.get("SMTP_PASSWORD", "")


def send_email(to_address, subject, body):
    if not SMTP_USER or not SMTP_PASSWORD:
        print(f"[email skipped - no credentials] to={to_address} subject={subject} body={body}")
        return

    msg = MIMEText(body)
    msg["Subject"] = subject
    msg["From"] = SMTP_USER
    msg["To"] = to_address

    with smtplib.SMTP(SMTP_HOST, SMTP_PORT) as server:
        server.starttls()
        server.login(SMTP_USER, SMTP_PASSWORD)
        server.sendmail(SMTP_USER, [to_address], msg.as_string())


@app.route("/health")
def health():
    return "ok"


@app.route("/internal/notifications/queue", methods=["POST"])
def queue_notification():
    data = request.get_json()

    user_id = data.get("user_id")
    event_type = data.get("event_type")
    payload = data.get("payload")

    print(f"[notification] to user {user_id}: {event_type} — {payload}")

    to_address = payload.get("email", "unknown@example.com")
    subject = "SellEasy notification"
    body = f"You have a new SellEasy notification: {event_type} — {payload}"

    send_email(to_address, subject, body)

    return jsonify({"queued": True})


if __name__ == "__main__":
    app.run(port=8000)