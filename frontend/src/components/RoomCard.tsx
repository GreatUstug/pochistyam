import type { Room } from '../types'

type Props = {
    room: Room
    onClean: (id: number) => void
    cleaning?: boolean
}

export function RoomCard({ room, onClean, cleaning }: Props) {
    const color =
        room.dirtiness >= 0.8 ? '#e53935' :
            room.dirtiness >= 0.5 ? '#fb8c00' :
                '#43a047'

    const percent = Math.round(room.dirtiness * 100)

    return (
        <div style={{
            border: '1px solid #ddd',
            borderRadius: 12,
            padding: 16,
            marginBottom: 12,
            fontFamily: 'sans-serif',
            maxWidth: 360,
        }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <h3 style={{ margin: 0 }}>{room.name}</h3>
                {room.needsCleaning && (
                    <span style={{
                        background: '#e53935', color: '#fff',
                        fontSize: 12, padding: '2px 8px', borderRadius: 8,
                    }}>
            пора
          </span>
                )}
            </div>

            <div style={{ margin: '12px 0' }}>
                <div style={{ height: 8, background: '#eee', borderRadius: 4, overflow: 'hidden' }}>
                    <div style={{
                        width: `${percent}%`,
                        height: '100%',
                        background: color,
                        transition: 'width 0.3s ease',
                    }} />
                </div>
                <small style={{ color: '#666' }}>Загрязнение: {percent}%</small>
            </div>

            <button
                onClick={() => onClean(room.id)}
                disabled={cleaning}
                style={{
                    padding: '8px 16px',
                    borderRadius: 8,
                    border: 'none',
                    background: '#1e88e5',
                    color: '#fff',
                    cursor: cleaning ? 'wait' : 'pointer',
                }}
            >
                {cleaning ? 'Убираем...' : 'Убрано'}
            </button>
        </div>
    )
}