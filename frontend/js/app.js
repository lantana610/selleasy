const API_BASE = "http://localhost:8080/api/v1";

function renderListings(listings) {
  const container = document.getElementById("listings");
  container.innerHTML = "";

  for (const listing of listings) {
    if (listing.status !== "available") {
      continue;
    }

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
  if (!currentUser) {
    alert("Please log in first.");
    return;
  }

  showOrderModal(listing);
}
function showOrderModal(listing) {
  const modal = document.getElementById("orderModal");
  modal.innerHTML = `
    <div class="modal-box">
      <h3>Confirm order</h3>
      <p>${listing.title} — ${listing.currency} ${listing.price}</p>
      <button class="confirm-btn">Confirm</button>
      <button class="cancel-btn">Cancel</button>
    </div>
  `;

  modal.querySelector(".confirm-btn").addEventListener("click", () => {
    confirmOrder(listing);
  });

  modal.querySelector(".cancel-btn").addEventListener("click", () => {
    closeOrderModal();
  });

  modal.classList.add("visible");
}

function closeOrderModal() {
  const modal = document.getElementById("orderModal");
  modal.classList.remove("visible");
}
let currentUser = null;

function loadListings() {
  fetch(`${API_BASE}/listings`)
    .then(response => response.json())
    .then(listings => {
      renderListings(listings);
    })
    .catch(error => {
      console.error("Failed to fetch listings:", error);
    });
}
loadListings();
function confirmOrder(listing) {
  fetch(`${API_BASE}/orders`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      listing_id: listing.id,
      buyer_id: currentUser.id,
      payment_method: "mobile_money",
      amount: listing.price,
      currency: listing.currency,
    }),
  })
    .then(response => {
      if (!response.ok) {
        return response.text().then(message => {
          throw new Error(message);
        });
      }
      return response.json();
    })
    .then(order => {
      console.log("Order placed:", order);
      closeOrderModal();
      loadListings();
    })
    .catch(error => {
      console.error("Failed to place order:", error);
      alert(error.message);
      closeOrderModal();
    });
}

const loginButton = document.getElementById("loginButton");
loginButton.addEventListener("click", () => {
  const email = document.getElementById("loginEmail").value;
  const password = document.getElementById("loginPassword").value;

  fetch(`${API_BASE}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  })
    .then(response => {
      if (!response.ok) {
        throw new Error("Login failed");
      }
      return response.text();
    })
    .then(result => {
      const userID = result.split(": ")[1];
      currentUser = { email: email, id: userID };
      document.getElementById("loginStatus").textContent = `Logged in as ${email}`;
    })
    .catch(error => {
      document.getElementById("loginStatus").textContent = "Login failed";
      console.error(error);
    });
});
const followButton = document.getElementById("followButton");
followButton.addEventListener("click", () => {
  if (!currentUser) {
    alert("Please log in first.");
    return;
  }

  const category = document.getElementById("followCategory").value;
  const city = document.getElementById("followCity").value;

  fetch(`${API_BASE}/follows`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      user_id: currentUser.id,
      category: category,
      city: city,
    }),
  })
    .then(response => response.json())
    .then(follow => {
      console.log("Follow created:", follow);
      document.getElementById("followStatus").textContent = "Following!";
    })
    .catch(error => {
      console.error("Failed to create follow:", error);
      document.getElementById("followStatus").textContent = "Failed to follow";
    });
});

const myOrdersButton = document.getElementById("myOrdersButton");
myOrdersButton.addEventListener("click", () => {
  if (!currentUser) {
    alert("Please log in first.");
    return;
  }

  fetch(`${API_BASE}/orders?buyer_id=${currentUser.id}`)
    .then(response => response.json())
    .then(orders => {
      renderMyOrders(orders);
    })
    .catch(error => {
      console.error("Failed to fetch orders:", error);
    });
});

function renderMyOrders(orders) {
  const container = document.getElementById("myOrdersList");
  container.innerHTML = "";

  if (orders.length === 0) {
    container.innerHTML = "<p>You have no orders yet.</p>";
    return;
  }

  for (const order of orders) {
    const item = document.createElement("div");
    item.className = "order-item";
    item.innerHTML = `
      <p>${order.currency} ${order.amount} — <strong>${order.status}</strong></p>
    `;
    container.appendChild(item);
  }
}


