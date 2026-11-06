import type { Room, RoomRequest } from './types'

export async function getRoom(req: RoomRequest): Promise<Room> {
    const res = await fetch('/api/rooms/get', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(req),
    })
    if (!res.ok) throw new Error(`HTTP ${res.status}`)
    return res.json() as Promise<Room>
}