let currentPrompt = "$ ";
let username = "";
let host = "";
let pwd = "";

function getUserName() {
    return new Promise((resolve, reject) => {
      const usernameSocket = new WebSocket(`ws://${window.location.host}/username`);
      usernameSocket.onopen = () => {
        console.log("[INFO] Username WebSocket connected.");
      };
      usernameSocket.onmessage = event => {
        console.log("[DEBUG] Username received:", event.data);
        username = event.data;
        resolve(username);
      };
      usernameSocket.onerror = error => {
        console.error("[ERROR] Username WebSocket failed:", error);
        reject(new Error("Failed to fetch username."));
      };
      usernameSocket.onclose = () => {
        console.log("[INFO] Username WebSocket closed.");
      };
    });
  }
  
  function getHost() {
    return new Promise((resolve, reject) => {
      const hostSocket = new WebSocket(`ws://${window.location.host}/hostname`);
      hostSocket.onopen = () => {
        console.log("[INFO] Host WebSocket connected.");
      };
      hostSocket.onmessage = event => {
        console.log("[DEBUG] Host received:", event.data);
        host = event.data;
        resolve(host);
      };
      hostSocket.onerror = error => {
        console.error("[ERROR] Host WebSocket failed:", error);
        reject(new Error("Failed to fetch host."));
      };
      hostSocket.onclose = () => {
        console.log("[INFO] Host WebSocket closed.");
      };
    });
  }
  
  function getPWD() {
    return new Promise((resolve, reject) => {
      const pwdSocket = new WebSocket(`ws://${window.location.host}/dir`);
      pwdSocket.onopen = () => {
        console.log("[INFO] PWD WebSocket connected.");
      };
      pwdSocket.onmessage = event => {
        console.log("[DEBUG] PWD received:", event.data);
        pwd = event.data;
        resolve(pwd);
      };
      pwdSocket.onerror = error => {
        console.error("[ERROR] PWD WebSocket failed:", error);
        reject(new Error("Failed to fetch PWD."));
      };
      pwdSocket.onclose = () => {
        console.log("[INFO] PWD WebSocket closed.");
      };
    });
  }
  