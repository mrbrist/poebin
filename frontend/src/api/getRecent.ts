import { API_BASE } from "./consts";

export interface buildData {
    ID: string;
    Level: number;
    Class: string;
    DateAdded: number;
}
export interface buildDataArr extends Array<buildData> {}

export async function getRecent() {
    try {
        const res = await fetch(`${API_BASE}/recent`, {
            method: "GET",
        });
        if (!res.ok) throw new Error("Failed to get recent builds");
        const recent: buildDataArr = await res.json();
        return recent;
    } catch (err) {
        // Display error to user
        console.log(err);
        return null;
    }
}
