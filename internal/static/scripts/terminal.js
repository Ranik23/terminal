const term = new Terminal({
  cursorBlink: true,
  cursorStyle: 'block', 
  rows: 150,
  cols: 150,
});

term.onKey(function (e) {
  const key = e.key; 
  const ev = e.domEvent; 
  
  console.log(`Нажата клавиша: ${key}`);

  if (key === "ArrowDown" || key === "ArrowUp" || key === "ArrowLeft" || key === "ArrowRight" ||
      key === "Enter" || key === "Backspace" || key === "Tab" || key === "Esc") {
    ev.preventDefault();
    console.log(`Заблокирована клавиша: ${key}`);
  }
});

let inputBuffer = "";

term.open(document.getElementById("terminal"));

term.writeln("\x1b[32mWelcome to the Web Terminal!\x1b[0m");
term.writeln("Type your commands below:");

createWebSocket(term, () => updatePrompt(term));

term.onData(data => {
  switch (data) {
    case "\r": 
      handleEnter(term);
      break;
    case "\u007f": 
      handleBackspace(term);
      break;
    default:
      handleDefault(data);
      break;
  }
});

function handleEnter(term) {
  if (inputBuffer.trim() !== "") {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(inputBuffer);
    } else {
      term.write("\r\n\x1b[31m[ERROR] WebSocket not connected\x1b[0m");
    }
  }
  inputBuffer = ""; 
}

function handleBackspace(term) {
  if (inputBuffer.length > 0) {
    inputBuffer = inputBuffer.slice(0, -1); 
    term.write("\b \b"); 
  }
}

function handleDefault(data) {
  inputBuffer += data; 
  term.write(data);
}

function updatePrompt(term) {
  const prompt = `${username}@${host}:${pwd}$ `;
  term.write(`\r\n${prompt}`);
}
