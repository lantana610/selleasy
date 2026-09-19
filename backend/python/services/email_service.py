import os
import smtplib
from email.mime.text import MIMEText

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