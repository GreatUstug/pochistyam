import { useEffect, useState } from 'react';
import type { Room } from './types';
import { getAllRooms, createRoom } from './api';
import { RoomCard } from './components/RoomCard';
import './App.css';

function App() {
    const [rooms, setRooms] = useState<Room[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const [name, setName] = useState('');
    const [width, setWidth] = useState(2);
    const [height, setHeight] = useState(2);
    const [isSubmitting, setIsSubmitting] = useState(false);

    const loadRooms = async () => {
        try {
            setLoading(true);
            setError(null);
            const data = await getAllRooms();
            setRooms(data);
        } catch (err: any) {
            setError(err.message || 'Не удалось загрузить комнаты');
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadRooms();
    }, []);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!name.trim()) return;

        setIsSubmitting(true);
        try {
            await createRoom({ name, width, height });
            setName('');
            setWidth(2);
            setHeight(2);
            await loadRooms();
        } catch (err: any) {
            setError(err.message || 'Ошибка при создании');
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="container">
            <header>
                <h1>Почистям</h1>
                <p>Менеджер уборки</p>
            </header>

            <section className="create-form">
                <h2>Добавить комнату</h2>
                <form onSubmit={handleSubmit} style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
                    <input
                        type="text"
                        placeholder="Название (напр. Ванная)"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        required
                    />
                    <input
                        type="number"
                        placeholder="Ширина"
                        value={width}
                        onChange={(e) => setWidth(Number(e.target.value))}
                        min="1" required
                    />
                    <input
                        type="number"
                        placeholder="Высота"
                        value={height}
                        onChange={(e) => setHeight(Number(e.target.value))}
                        min="1" required
                    />
                    <button type="submit" disabled={isSubmitting}>
                        {isSubmitting ? 'Создание...' : 'Создать'}
                    </button>
                </form>
            </section>

            <section className="rooms-list">
                <h2>Мои комнаты</h2>

                {loading && <p>Загрузка...</p>}
                {error && <p className="error">{error}</p>}

                {!loading && rooms.length === 0 && (
                    <p style={{ color: '#666' }}>У вас пока нет комнат. Создайте первую!</p>
                )}

                {rooms.map((room) => (
                    <RoomCard key={room.id} room={room} />
                ))}
            </section>
        </div>
    );
}

export default App;