package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "First time when pochistyam is working")
		if err != nil {
			return
		}
	})
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

const indexHTML = `<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="utf-8">
  <title>По чистям</title>
</head>
<body>
  <h1>Счетчик нажатий</h1>
  <button id="btn">Нажми меня</button>
  <p>Нажато: <span id="count">0</span></p>

  <script>
    const btn = document.getElementById('btn');
    const countEl = document.getElementById('count');

    // При загрузке страницы подтягиваем текущее значение
    fetch('/count')
      .then(r => r.json())
      .then(data => { countEl.textContent = data.count; })
      .catch(err => console.error(err));

    btn.addEventListener('click', async () => {
      try {
        const res = await fetch('/increment', { method: 'POST' });
        if (!res.ok) throw new Error('HTTP ' + res.status);
        const data = await res.json();
        countEl.textContent = data.count;
      } catch (err) {
        console.error('Ошибка:', err);
        countEl.textContent = 'ошибка';
      }
    });
  </script>
</body>
</html>`
