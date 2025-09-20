document.addEventListener("DOMContentLoaded", () => {
    const fetchBtn = document.querySelector("#fetchbutton");
    const responseEl = document.querySelector("#response");
    fetchBtn.addEventListener("click", async () => {
        try {
            const response = await fetch("/api/joke");
            const data = await response.json();
            responseEl.textContent = data.message;
        } catch (error) {
            responseEl.textContent = `An error occured ${error}`;
        }
    })
})