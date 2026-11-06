import { useState } from 'react'
import type { Room } from './types'
import { getRoom } from './api'
import { RoomCard } from './components/RoomCard'

function App() {
    const [room, setRoom] = useState<Room | null>(null)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const handleLoad = async () => {
        setLoading(true)
        setError(null)
        try {
            const data = await getRoom({ clientId: 'client-uuid-123', roomId: 1 })
            setRoom(data)
        } catch {
            setError('Не удалось загрузить комнату')
        } finally {
            setLoading(false)
        }
    }

    return (
        <div style={{ padding: 40, fontFamily: 'consolas', maxWidth: 480 }}>
            <h1>Почистям</h1>
            <button onClick={handleLoad} disabled={loading}>
                {loading ? 'Загрузка...' : 'Загрузить комнату'}
            </button>

            {error && <p style={{ color: 'red' }}>{error}</p>}

            {room && (
                <div style={{ marginTop: 20 }}>
                    <RoomCard room={room} onClean={() => {}} />
                </div>
            )}
        </div>
    )
}

export default App