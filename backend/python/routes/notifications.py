from flask import request, jsonify

from services.email_service import send_email


def register_routes(app):
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
