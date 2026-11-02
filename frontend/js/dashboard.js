const gate = document.getElementById("loginGate");
const content = document.getElementById("dashboardContent");

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

function showDashboard() {
  const session = getSession();
  gate.hidden = Boolean(session);
  content.hidden = !session;
  if (session) {
    document.getElementById("welcome").textContent = "Welcome, " + session.email;
    loadMyListings();
  }
}

window.addEventListener("session-changed", showDashboard);
showDashboard();