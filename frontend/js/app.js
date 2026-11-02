const listingsEl = document.getElementById("listings");
let shownListings = [];

async function loadListings() {
  try {
    shownListings = await api("/listings");
    drawListings(shownListings);
  } catch (err) {
    listingsEl.textContent = err.message;
  }
}

function drawListings(listings) {
  if (listings.length === 0) {
    listingsEl.innerHTML = "<p>No items yet.</p>";
    return;
  }

  listingsEl.innerHTML = listings.map(function (listing) {
    return (
      '<article class="product-card">' +
      "<h3>" + esc(listing.title) + "</h3>" +
      "<p>" + esc(listing.currency) + " " + esc(listing.price) + "</p>" +
      "<p>" + esc(listing.city) + ", " + esc(listing.country) + "</p>" +
      '<button type="button" data-id="' + esc(listing.id) + '">Order</button>' +
      "</article>"
    );
  }).join("");
}

listingsEl.addEventListener("click", function (event) {
  const button = event.target.closest("button[data-id]");
  if (!button) return;

  const listing = shownListings.find(function (item) {
    return item.id === button.dataset.id;
  });
  startOrder(listing);
});

function startOrder(listing) {
  const message = document.getElementById("pageMessage");
  if (!getSession()) {
    setMessage(message, "Log in to order.", "error");
    return;
  }
  openOrderDialog(listing);
}

function openOrderDialog(listing) {
  const modal = document.getElementById("orderModal");
  modal.innerHTML =
    '<div class="modal-box">' +
    "<h3>Confirm order</h3>" +
    "<p>" + esc(listing.title) + " - " + esc(listing.currency) + " " + esc(listing.price) + "</p>" +
    '<label for="paymentMethod">How will you pay?</label>' +
    '<select id="paymentMethod">' +
    '<option value="cash_on_delivery">Cash on delivery</option>' +
    '<option value="mobile_money">Mobile money</option>' +
    '<option value="bank_transfer">Bank transfer</option>' +
    '<option value="pickup">Pickup</option>' +
    "</select>" +
    '<p id="orderMessage" class="message"></p>' +
    '<button type="button" id="confirmOrder">Confirm order</button> ' +
    '<button type="button" id="cancelOrder">Cancel</button>' +
    "</div>";
  modal.classList.add("visible");

  document.getElementById("cancelOrder").addEventListener("click", closeOrderDialog);
  document.getElementById("confirmOrder").addEventListener("click", function () {
    sendOrder(listing);
  });
}

function closeOrderDialog() {
  const modal = document.getElementById("orderModal");
  modal.classList.remove("visible");
  modal.innerHTML = "";
}
async function sendOrder(listing) {
  const orderMessage = document.getElementById("orderMessage");
  const confirmButton = document.getElementById("confirmOrder");
  confirmButton.disabled = true;

  try {
    await api("/orders", {
      method: "POST",
      body: {
        listing_id: listing.id,
        payment_method: document.getElementById("paymentMethod").value,
      },
    });
    closeOrderDialog();
    setMessage(document.getElementById("pageMessage"), "Order sent. The seller has been notified.", "ok");
    loadListings();
  } catch (err) {
    setMessage(orderMessage, err.message, "error");
    confirmButton.disabled = false;
  }
}
loadListings();