function login() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const messageDiv = document.getElementById("message");

    fetch("http://localhost:8080/api/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
    })
        .then(response => response.json())
        .then(data => {

            if (data.status === 200 && data.data.success) {
                messageDiv.style.color = "green";
                messageDiv.innerText = "Login berhasil!";
                localStorage.setItem("token", data.data.token); // Simpan token JWT
                window.location.href = "chat.html"; // Redirect ke halaman chat
            } else {
                messageDiv.style.color = "red";
                messageDiv.innerText = "Login gagal: " + data.message;
            }
        })
        .catch(error => {
            messageDiv.style.color = "red";
            messageDiv.innerText = "Terjadi kesalahan, coba lagi.";
        });
}
