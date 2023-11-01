let btn = document.getElementById("login-btn");
btn.addEventListener("click", (event) => {
  btn.classList.toggle("is-loading");
});
document.getElementById("loginform").addEventListener(
  "htmx:afterOnLoad",
  (event) => {
    btn.classList.toggle("is-loading");
    console.log("toggled");
  },
  (useCapture = false)
);
