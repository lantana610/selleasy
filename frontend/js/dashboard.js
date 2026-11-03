const gate = document.getElementById("loginGate");
const content = document.getElementById("dashboardContent");

const PAYMENT_LABELS = {
  cash_on_delivery: "Cash on delivery",
  mobile_money: "Mobile money",
  bank_transfer: "Bank transfer",
  pickup: "Pickup",
};

async function loadMyListings() {
  const list = document.getElementById("myListings");
  try {
    const listings = await api("/listings/mine");
    if (listings.length === 0) {
      list.innerHTML = "<li>You haven't listed anything yet.</li>";
      return;
    }
    list.innerHTML = listings.map(function (listing) {
      return (
        "<li>" + esc(listing.title) + " - " + esc(listing.currency) + " " +
        esc(listing.price) + " (" + esc(listing.status) + ")</li>"
      );
    }).join("");
  } catch (err) {
    setMessage(document.getElementById("dashMessage"), err.message, "error");
  }
}

async function loadSellerOrders() {
  const list = document.getElementById("sellerOrders");
  try {
    const orders = await api("/orders?role=seller");
    if (orders.length === 0) {
      list.innerHTML = "<li>No orders on your listings yet.</li>";
      return;
    }
    list.innerHTML = orders.map(function (order) {
      const buttons = order.status === "requested"
        ? '<br><button type="button" data-order="' + esc(order.id) + '" data-status="completed">Mark as sold</button> ' +
          '<button type="button" data-order="' + esc(order.id) + '" data-status="cancelled">Decline</button>'
        : "";
      return (
        "<li>" + esc(order.listing_title) + " - " + esc(order.currency) + " " +
        esc(order.amount) + " (" + esc(order.status) + ")<br>" +
        "Buyer: " + esc(order.counterpart_name) + ", " + esc(order.counterpart_email) +
        ", " + esc(order.counterpart_phone) + "<br>" +
        "Payment: " + esc(PAYMENT_LABELS[order.payment_method] || order.payment_method) +
        buttons + "</li>"
      );
    }).join("");
  } catch (err) {
    setMessage(document.getElementById("dashMessage"), err.message, "error");
  }
}

document.getElementById("sellerOrders").addEventListener("click", async function (event) {
  const button = event.target.closest("button[data-order]");
  if (!button) return;

  const status = button.dataset.status;
  if (status === "cancelled" && !confirm("Decline this order?")) return;

  button.disabled = true;
  try {
    await api("/orders/" + button.dataset.order + "/status", {
      method: "PATCH",
      body: { status: status },
    });
    setMessage(document.getElementById("dashMessage"),
      status === "completed" ? "Marked as sold." : "Order declined.", "ok");
  } catch (err) {
    setMessage(document.getElementById("dashMessage"), err.message, "error");
  }
  loadMyListings();
  loadSellerOrders();
});
function showDashboard() {
  const session = getSession();
  gate.hidden = Boolean(session);
  content.hidden = !session;
  if (session) {
    document.getElementById("welcome").textContent = "Welcome, " + session.email;
    loadMyListings();
    loadSellerOrders();
  }
}

window.addEventListener("session-changed", showDashboard);
showDashboard();