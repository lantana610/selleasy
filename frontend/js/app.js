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
  setMessage(message, "You chose: " + listing.title, "ok");
}

loadListings();