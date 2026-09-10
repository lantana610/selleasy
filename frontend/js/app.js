const API_BASE = "http://localhost:8080/api/v1";

fetch(`${API_BASE}/listings`)
  .then(response => response.json())
  .then(listings => {
    renderListings(listings);
  })
  .catch(error => {
    console.error("Failed to fetch listings:", error);
  });

function renderListings(listings) {
  const container = document.getElementById("listings");

  for (const listing of listings) {
    const card = document.createElement("div");
    card.className = "product-card";
    card.innerHTML = `
      <h3>${listing.title}</h3>
      <p>${listing.currency} ${listing.price}</p>
      <p>${listing.city}, ${listing.country}</p>
      <button>Order</button>
    `;

    const button = card.querySelector("button");
    button.addEventListener("click", () => {
      placeOrder(listing);
    });

    container.appendChild(card);
  }
}

function placeOrder(listing) {
  fetch(`${API_BASE}/orders`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      listing_id: listing.id,
      buyer_id: "b1",
      payment_method: "mobile_money",
      amount: listing.price,
      currency: listing.currency,
    }),
  })
    .then(response => response.json())
    .then(order => {
      console.log("Order placed:", order);
      alert(`Order placed for ${listing.title}!`);
    })
    .catch(error => {
      console.error("Failed to place order:", error);
    });
}