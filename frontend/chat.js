const token = localStorage.getItem("token");
if (!token) {
    window.location.href = "login.html"; // Redirect jika belum login
}

const ws = new WebSocket("ws://localhost:8080/ws");
let replyTo = null; // Menyimpan ID pesan yang akan dibalas

ws.onopen = () => {
    console.log("Terhubung ke WebSocket");
};

ws.onmessage = (event) => {
    const messageDiv = document.getElementById("messages");
    const data = JSON.parse(event.data);

    const messageElement = document.createElement("div");
    messageElement.classList.add("message");

    if (data.replyTo) {
        messageElement.innerHTML = `<i>Balasan ke #${data.replyTo}:</i> ${data.message}`;
    } else {
        messageElement.innerHTML = `#${data.id}: ${data.message}`;
    }

    // Tombol "Balas"
    const replyBtn = document.createElement("span");
    replyBtn.textContent = " Balas";
    replyBtn.classList.add("reply");
    replyBtn.onclick = () => setReplyTo(data.id);

    messageElement.appendChild(replyBtn);
    messageDiv.appendChild(messageElement);
};

ws.onclose = () => {
    console.log("Koneksi WebSocket tertutup.");
};

function sendMessage() {
    const messageInput = document.getElementById("messageInput");
    const message = messageInput.value;

    if (message.trim() !== "") {
        ws.send(JSON.stringify({ message, replyTo }));
        messageInput.value = "";
        replyTo = null;
    }
}

function setReplyTo(id) {
    replyTo = id;
    document.getElementById("messageInput").placeholder = `Balas pesan #${id}...`;
}

function logout() {
    localStorage.removeItem("token");
    window.location.href = "login.html";
}
