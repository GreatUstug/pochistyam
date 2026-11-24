import { useEffect, useState } from 'react';
import type { Room } from './types';
import { getAllRooms, createRoom } from './api';
import { RoomCard } from './components/RoomCard';
import { AuthForm } from './components/AuthForm';
import { isAuthenticated, removeToken } from './auth';
import './App.css';

function App() {
    const [isAuth, setIsAuth] = useState(isAuthenticated());
    const [rooms, setRooms] = useState<Room[]>([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const [name, setName] = useState('');
    const [width, setWidth] = useState(2);
    const [height, setHeight] = useState(2);
    const [isSubmitting, setIsSubmitting] = useState(false);

    const loadRooms = async () => {
        if (!isAuth) return;
        try {
            setLoading(true);
            setError(null);
            const data = await getAllRooms();
            setRooms(data);
        } catch (err: any) {
            if (err.message === 'UNAUTHORIZED') {
                handleLogout();
            } else {
                setError(err.message || 'Ошибка загрузки');
            }
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        if (isAuth) loadRooms();
    }, [isAuth]);

    const handleLogout = () => {
        removeToken();
        setIsAuth(false);
        setRooms([]);
    };

    const handleCreate = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsSubmitting(true);
        try {
            await createRoom({ name, width, height });
            setName('');
            await loadRooms();
        } catch (err: any) {
            setError(err.message);
        } finally {
            setIsSubmitting(false);
        }
    };

    if (!isAuth) {
        return <AuthForm onSuccess={() => setIsAuth(true)} />;
    }

    return (
        <div className="container">
            <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <div>
                    <h1>Почистям</h1>
                    <p>Менеджер уборки</p>
                </div>
                <button onClick={handleLogout} style={{ background: '#e74c3c', height: 'fit-content' }}>
                    Выйти
                </button>
            </header>

            <section className="create-form">
                <h2>Добавить комнату</h2>
                <form onSubmit={handleCreate} style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
                    <input type="text" placeholder="Название" value={name} onChange={e => setName(e.target.value)} required />
                    <input type="number" placeholder="Ширина" value={width} onChange={e => setWidth(Number(e.target.value))} min="1" required />
                    <input type="number" placeholder="Высота" value={height} onChange={e => setHeight(Number(e.target.value))} min="1" required />
                    <button type="submit" disabled={isSubmitting}>
                        {isSubmitting ? 'Создание...' : 'Создать'}
                    </button>
                </form>
            </section>

            <section className="rooms-list">
                <h2>Мои комнаты</h2>
                {loading && <p>Загрузка...</p>}
                {error && <p className="error">{error}</p>}
                {!loading && rooms.length === 0 && <p style={{ color: '#666' }}>У вас пока нет комнат.</p>}
                {rooms.map(room => (
                    <RoomCard key={room.id} room={room} />
                ))}
            </section>
        </div>
    );
}

export default App;