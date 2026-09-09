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
    `;
    container.appendChild(card);
  }
}