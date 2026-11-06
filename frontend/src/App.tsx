import { useEffect, useState } from 'react'

type CountResponse = {
  count: number
}

function App() {
  const [count, setCount] = useState<number>(0)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetch('/api/count')
        .then((r) => r.json() as Promise<CountResponse>)
        .then((data) => setCount(data.count))
        .catch(() => setError('Не удалось загрузить счётчик'))
  }, [])

  const handleClick = async () => {
    setLoading(true)
    setError(null)
    try {
      const res = await fetch('/api/increment', { method: 'POST' })
      if (!res.ok) throw new Error(`HTTP ${res.status}`)
      const data = (await res.json()) as CountResponse
      setCount(data.count)
    } catch {
      setError('Ошибка запроса')
    } finally {
      setLoading(false)
    }
  }

  return (
      <div style={{ padding: 40, fontFamily: 'sans-serif' }}>
        <h1>По чистям</h1>
        <button
            onClick={handleClick}
            disabled={loading}
            style={{ fontSize: 18, padding: '8px 24px', cursor: 'pointer' }}
        >
          {loading ? '...' : 'Нажми меня'}
        </button>
        <p>Нажато: <strong>{count}</strong></p>
        {error && <p style={{ color: 'red' }}>{error}</p>}
      </div>
  )
}

export default App