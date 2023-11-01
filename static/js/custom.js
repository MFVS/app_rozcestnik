function logHtmxEvents(watchedElem) {
  // Define the htmx events you want to listen to
  const events = [
    "htmx:afterSwap", // Triggered after content has been swapped in
    "htmx:beforeSwap", // Triggered before content is swapped in
    "htmx:afterRequest", // Triggered after an AJAX request is made
    "htmx:beforeRequest", // Triggered before an AJAX request is made
    "htmx:configRequest", // Triggered before the request is configured
    "htmx:responseError", // Triggered on an error response
    "htmx:afterOnLoad", // Triggered after a successful response has been loaded
    "htmx:afterProcess", // Triggered after the response has been processed
    "htmx:beforeOnLoad", // Triggered before a successful response is loaded
    "htmx:beforeProcess", // Triggered before the response is processed
    "htmx:afterValidate", // Triggered after input validation
    "htmx:beforeValidate", // Triggered before input validation
    "htmx:validationFailed", // Triggered when input validation fails
  ];

  // Loop through the events and add event listeners
  events.forEach((event) => {
    watchedElem.addEventListener(
      event,
      (e) => {
          console.log(`HTMX Event: ${event} ${e.target}`);
      },
      (useCapture = false)
    );
  });
}

// Call the function to set up event listeners
logHtmxEvents(document.body);

// let btn = document.getElementById("login-btn");
// logHtmxEvents(btn);
