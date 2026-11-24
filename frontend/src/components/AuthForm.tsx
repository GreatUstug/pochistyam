import { useState } from 'react';
import { loginUser, registerUser } from '../api';
import { saveToken } from '../auth';

type Props = {
    onSuccess: () => void;
};

export function AuthForm({ onSuccess }: Props) {
    const [isLogin, setIsLogin] = useState(true);
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError(null);
        setLoading(true);

        try {
            let data;
            if (isLogin) {
                data = await loginUser(email, password);
            } else {
                data = await registerUser(name, email, password);
            }
            saveToken(data.token);
            onSuccess();
        } catch (err: any) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div style={{ maxWidth: 400, margin: '40px auto', padding: 20, background: '#fff', borderRadius: 8, boxShadow: '0 2px 8px rgba(0,0,0,0.1)' }}>
            <h2 style={{ textAlign: 'center' }}>{isLogin ? 'Вход' : 'Регистрация'}</h2>

            <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: 10 }}>
                {!isLogin && (
                    <input
                        type="text" placeholder="Имя" value={name}
                        onChange={e => setName(e.target.value)} required
                        style={{ padding: 8, borderRadius: 4, border: '1px solid #ddd' }}
                    />
                )}
                <input
                    type="email" placeholder="Email" value={email}
                    onChange={e => setEmail(e.target.value)} required
                    style={{ padding: 8, borderRadius: 4, border: '1px solid #ddd' }}
                />
                <input
                    type="password" placeholder="Пароль" value={password}
                    onChange={e => setPassword(e.target.value)} required
                    style={{ padding: 8, borderRadius: 4, border: '1px solid #ddd' }}
                />

                {error && <p style={{ color: 'red', fontSize: 14, margin: 0 }}>{error}</p>}

                <button type="submit" disabled={loading} style={{ padding: 10, background: '#3498db', color: '#fff', border: 'none', borderRadius: 4, cursor: 'pointer' }}>
                    {loading ? 'Загрузка...' : (isLogin ? 'Войти' : 'Зарегистрироваться')}
                </button>
            </form>

            <p style={{ textAlign: 'center', marginTop: 15, fontSize: 14 }}>
                {isLogin ? 'Нет аккаунта? ' : 'Уже есть аккаунт? '}
                <button
                    onClick={() => { setIsLogin(!isLogin); setError(null); }}
                    style={{ background: 'none', border: 'none', color: '#3498db', cursor: 'pointer', padding: 0, textDecoration: 'underline' }}
                >
                    {isLogin ? 'Зарегистрироваться' : 'Войти'}
                </button>
            </p>
        </div>
    );
}