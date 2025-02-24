const API_URL = "http://localhost:8080";  // Ganti sesuai URL backend-mu

// Login user
async function login() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const response = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
    });

    const data = await response.json();
    if (response.ok) {
        localStorage.setItem("token", data.token);
        document.getElementById("login-container").style.display = "none";
        document.getElementById("chat-container").style.display = "block";
        loadMessages();
    } else {
        document.getElementById("login-error").innerText = data.error;
    }
}

// Logout user
function logout() {
    localStorage.removeItem("token");
    document.getElementById("login-container").style.display = "block";
    document.getElementById("chat-container").style.display = "none";
}

// Mengambil pesan dari server
async function loadMessages() {
    const token = localStorage.getItem("token");
    if (!token) return logout();

    const response = await fetch(`${API_URL}/messages`, {
        headers: { "Authorization": token }
    });

    const messages = await response.json();
    const chatBox = document.getElementById("chat-box");
    chatBox.innerHTML = messages.map(msg => `<p><b>${msg.sender}:</b> ${msg.text}</p>`).join("");
}

// Mengirim pesan ke server
async function sendMessage() {
    const token = localStorage.getItem("token");
    if (!token) return logout();

    const message = document.getElementById("message").value;
    if (!message) return;

    await fetch(`${API_URL}/messages`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token
        },
        body: JSON.stringify({ text: message })
    });

    document.getElementById("message").value = "";
    loadMessages();
}

// Load pesan setiap 3 detik
setInterval(loadMessages, 3000);
