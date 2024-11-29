const term = new Terminal({
  cursorBlink: true, // Для мигающего курсора
  cursorStyle: 'block', // Для блок-курсиора
  rows: 150,
  cols: 150,
});

term.onKey(function (e) {
  const key = e.key; // Получаем нажатую клавишу
  const ev = e.domEvent; // Получаем оригинальное событие
  
  console.log(`Нажата клавиша: ${key}`);

  if (key === "ArrowDown" || key === "ArrowUp" || key === "ArrowLeft" || key === "ArrowRight" ||
      key === "Enter" || key === "Backspace" || key === "Tab" || key === "Esc") {
    ev.preventDefault(); // Блокируем прокрутку, выполнение действия по умолчанию
    console.log(`Заблокирована клавиша: ${key}`);
  }
});

// Буфер для ввода текста
let inputBuffer = "";

// Открываем терминал в указанном элементе DOM
term.open(document.getElementById("terminal"));

// Приветственное сообщение
term.writeln("\x1b[32mWelcome to the Web Terminal!\x1b[0m");
term.writeln("Type your commands below:");

// Устанавливаем WebSocket-соединение
createWebSocket(term, () => updatePrompt(term));

// Обработка данных из ввода терминала
term.onData(data => {
  switch (data) {
    case "\r": // Enter
      handleEnter(term);
      break;
    case "\u007f": // Backspace
      handleBackspace(term);
      break;
    default:
      handleDefault(data);
      break;
  }
});

// Обработка нажатия Enter
function handleEnter(term) {
  if (inputBuffer.trim() !== "") {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(inputBuffer); // Отправляем команду
    } else {
      term.write("\r\n\x1b[31m[ERROR] WebSocket not connected\x1b[0m");
    }
  }
  inputBuffer = ""; // Очищаем буфер после отправки
}

// Обработка Backspace
function handleBackspace(term) {
  if (inputBuffer.length > 0) {
    inputBuffer = inputBuffer.slice(0, -1); // Удаляем последний символ
    term.write("\b \b"); // Удаляем символ в терминале
  }
}

// Обработка ввода по умолчанию (любой другой символ)
function handleDefault(data) {
  inputBuffer += data; // Добавляем данные в буфер
  term.write(data); // Отображаем введенные данные в терминале
}

// Функция для обновления приглашения
function updatePrompt(term) {
  const prompt = `${username}@${host}:${pwd}$ `;
  term.write(`\r\n${prompt}`);
}
