const API_BASE = "http://localhost:8080/api/v1";

fetch(`${API_BASE}/listings`)
  .then(response => response.json())
  .then(listings => {
    console.log(listings);
  })
  .catch(error => {
    console.error("Failed to fetch listings:", error);
  });