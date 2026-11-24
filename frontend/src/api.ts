import type { CreateRoomRequest, Room, RoomsListResponse } from './types';
import { getToken } from './auth';

const BASE_URL = '/api';

function getHeaders() {
    const token = getToken();
    return {
        'Content-Type': 'application/json',
        ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
    };
}

export async function getAllRooms(): Promise<Room[]> {
    const res = await fetch(`${BASE_URL}/rooms`, { headers: getHeaders() });
    if (res.status === 401) throw new Error('UNAUTHORIZED'); // Чтобы фронт понял, что надо разлогиниться
    if (!res.ok) throw new Error(`Failed to fetch rooms: ${res.statusText}`);
    const data: RoomsListResponse = await res.json();
    return data.rooms || [];
}

export async function createRoom(req: CreateRoomRequest): Promise<Room> {
    const res = await fetch(`${BASE_URL}/createRoom`, {
        method: 'POST',
        headers: getHeaders(),
        body: JSON.stringify(req),
    });
    if (res.status === 401) throw new Error('UNAUTHORIZED');
    if (!res.ok) {
        const errorText = await res.text();
        throw new Error(errorText || `Failed to create room: ${res.statusText}`);
    }
    return res.json();
}

export async function loginUser(email: string, password: string) {
    const res = await fetch(`${BASE_URL}/auth`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
    });
    if (!res.ok) throw new Error('Неверный email или пароль');
    return res.json();
}

export async function registerUser(name: string, email: string, password: string) {
    const res = await fetch(`${BASE_URL}/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password }),
    });
    if (!res.ok) throw new Error('Ошибка регистрации (возможно, email занят)');
    return res.json();
}