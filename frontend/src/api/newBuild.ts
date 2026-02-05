import { API_BASE } from "./consts";

export interface newBuild {
    id: string;
}

export async function newBuild(raw: string) {
    try {
        const res = await fetch(`${API_BASE}/newBuild`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                raw: raw,
            }),
        });
        if (!res.ok) throw new Error("Failed to create build");
        const id: newBuild = await res.json();
        return id.id;
    } catch (err) {
        // Display error to user that the build is invalid
        console.log(err);
        return null;
    }
}
