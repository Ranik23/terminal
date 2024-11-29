let socket;

function createWebSocket(term, updatePrompt) {
  socket = new WebSocket(`ws://${window.location.host}/ws`);

  socket.onopen = () => {
    term.write("\r\n\x1b[32m[INFO] Connected to server\x1b[0m");
  
    Promise.all([getHost(), getUserName(), getPWD()])
      .then(() => {
        console.log("[INFO] All data fetched successfully.");
        updatePrompt(term);
      })
      .catch(error => {
        term.write("\r\n\x1b[31m[ERROR] Failed to fetch prompt data\x1b[0m");
        console.error("[ERROR] Error during WebSocket calls:", error);
      });
  };

  socket.onmessage = event => {
    if (event.data != "") {
      term.write("\r\n" + event.data);
    }

    Promise.all([getHost(), getUserName(), getPWD()])
      .then(() => {
        console.log("[INFO] All data fetched successfully.");
        updatePrompt(term);
      })
      .catch(error => {
        term.write("\r\n\x1b[31m[ERROR] Failed to fetch prompt data\x1b[0m");
        console.error("[ERROR] Error during WebSocket calls:", error);
      });
  };

  socket.onclose = () => {
    term.write("\r\n\x1b[31m[INFO] Connection closed. Retrying in 2 seconds...\x1b[0m");
    setTimeout(() => createWebSocket(term, updatePrompt), 2000);
  };

  socket.onerror = () => {
    term.write("\r\n\x1b[31m[ERROR] WebSocket error occurred\x1b[0m");
  };
}
