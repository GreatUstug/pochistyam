import type { CreateRoomRequest, Room, RoomsListResponse } from './types';

const BASE_URL = '/api';

export async function getAllRooms(): Promise<Room[]> {
    const res = await fetch(`${BASE_URL}/rooms`);
    if (!res.ok) {
        throw new Error(`Failed to fetch rooms: ${res.statusText}`);
    }
    const data: RoomsListResponse = await res.json();
    return data.rooms;
}

export async function createRoom(req: CreateRoomRequest): Promise<Room> {
    const res = await fetch(`${BASE_URL}/createRoom`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(req),
    });

    if (!res.ok) {
        const errorText = await res.text();
        throw new Error(errorText || `Failed to create room: ${res.statusText}`);
    }

    return res.json();
}