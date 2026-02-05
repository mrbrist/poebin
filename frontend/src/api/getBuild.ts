import { API_BASE } from "./consts";

export interface buildData {
    LastModified: number;
    Id: string;
    Raw: string;
    Data: JSON;
}

export async function getBuild(id: string) {
    try {
        const res = await fetch(`${API_BASE}/getBuild/${id}`, {
            method: "GET",
        });
        if (!res.ok) throw new Error("Failed to get build");
        const build: buildData = await res.json();
        return build;
    } catch (err) {
        // Display error to user
        console.log(err);
        return null;
    }
}
