const listingsEl = document.getElementById("listings");

async function loadListings() {
  try {
    const listings = await api("/listings");
    drawListings(listings);
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
      "</article>"
    );
  }).join("");
}

loadListings();