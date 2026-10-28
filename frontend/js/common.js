const API_BASE = "http://localhost:8080/api/v1";

localStorage.removeItem("currentUser");

function getSession() {
  try {
    return JSON.parse(localStorage.getItem("session"));
  } catch (e) {
    return null;
  }
}

function saveSession(session) {
  localStorage.setItem("session", JSON.stringify(session));
}

function clearSession() {
  localStorage.removeItem("session");
}

async function api(path, options = {}) {
  const headers = {};
  const session = getSession();
  if (session) {
    headers["Authorization"] = "Bearer " + session.token;
  }
  if (options.body !== undefined) {
    headers["Content-Type"] = "application/json";
  }

  const response = await fetch(API_BASE + path, {
    method: options.method || "GET",
    headers: headers,
    body: options.body === undefined ? undefined : JSON.stringify(options.body),
  });

  const text = await response.text();
  if (!response.ok) {
    throw new Error(text.trim() || "Something went wrong");
  }
  return text ? JSON.parse(text) : null;
}

function esc(value) {
  return String(value ?? "").replace(/[&<>"']/g, function (ch) {
    return { "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;" }[ch];
  });
}

function setMessage(element, text, kind) {
  if (!element) return;
  element.textContent = text || "";
  element.className = "message" + (kind ? " " + kind : "");
}

function announceSessionChange() {
  window.dispatchEvent(new Event("session-changed"));
}