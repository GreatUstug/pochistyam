import type { Room } from '../types';

type Props = {
    room: Room;
};

export function RoomCard({ room }: Props) {
    const dirtPercent = Math.min(Math.round(room.dirtiness * 100), 100);

    const barColor = dirtPercent >= 80 ? '#e53935' : dirtPercent >= 50 ? '#fb8c00' : '#43a047';

    return (
        <div style={{
            border: '1px solid #ddd',
            borderRadius: '8px',
            padding: '16px',
            marginBottom: '16px',
            backgroundColor: '#fff',
            boxShadow: '0 2px 4px rgba(0,0,0,0.05)'
        }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <h3 style={{ margin: 0, fontSize: '18px' }}>{room.name}</h3>
                {room.needsCleaning && (
                    <span style={{ backgroundColor: '#e53935', color: '#fff', padding: '2px 8px', borderRadius: '12px', fontSize: '12px' }}>
            Пора убирать!
          </span>
                )}
            </div>

            <div style={{ margin: '8px 0', fontSize: '14px', color: '#666' }}>
                Размер: {room.width} x {room.height}
            </div>

            <div style={{ marginTop: '12px' }}>
                <div style={{ height: '8px', backgroundColor: '#eee', borderRadius: '4px', overflow: 'hidden' }}>
                    <div style={{ width: `${dirtPercent}%`, height: '100%', backgroundColor: barColor, transition: 'width 0.3s' }} />
                </div>
                <div style={{ fontSize: '12px', color: '#666', marginTop: '4px' }}>
                    Загрязнение: {dirtPercent}%
                </div>
            </div>
        </div>
    );
}