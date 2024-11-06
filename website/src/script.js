// Function to fetch Lexer output
function fetchLexer() {
  fetch("https://cb.bynisarg.in/lexer")
    .then((response) => response.json())
    .then((data) => {
      displayOutput(data.output || data.error);
    })
    .catch((error) => {
      displayOutput("Error fetching lexer output: " + error.message);
    });
}

// Function to fetch Semantic output
function fetchSemantic() {
  fetch("https://cb.bynisarg.in/semantic")
    .then((response) => response.json())
    .then((data) => {
      displayOutput(data.output || data.error);
    })
    .catch((error) => {
      displayOutput("Error fetching semantic output: " + error.message);
    });
}

// Function to fetch 3AC output
function fetch3ac() {
  fetch("https://cb.bynisarg.in/3ac")
    .then((response) => response.json())
    .then((data) => {
      displayOutput(data.output || data.error);
    })
    .catch((error) => {
      displayOutput("Error fetching 3AC output: " + error.message);
    });
}

// Function to fetch all outputs
function fetchAll() {
  fetch("https://cb.bynisarg.in/all")
    .then((response) => response.json())
    .then((data) => {
      const output = `
                Lexer Output:\n${data.lexer}\n\n
                Semantic Output:\n${data.semantic}\n\n
                3AC Output:\n${data["3ac"]}\n
            `;
      displayOutput(output);
    })
    .catch((error) => {
      displayOutput("Error fetching all outputs: " + error.message);
    });
}

// Function to display the output in the output div
function displayOutput(output) {
  const outputDiv = document.getElementById("output");
  outputDiv.textContent = output;
}

// Function to upload a file
function uploadFile() {
  const fileInput = document.getElementById("fileUpload");
  const file = fileInput.files[0];

  if (!file) {
    displayOutput("Please select a file to upload.");
    return;
  }

  const formData = new FormData();
  formData.append("file", file);

  fetch("https://cb.bynisarg.in/upload", {
    method: "POST",
    body: formData,
  })
    .then((response) => response.json())
    .then((data) => {
      displayOutput(data.message || data.error);
    })
    .catch((error) => {
      displayOutput("Error uploading file: " + error.message);
    });
}
