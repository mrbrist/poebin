export function parseItem(raw: any) {
    raw = String(raw);
    raw = raw.replace(/\t/g, "");
    return raw.split("\n")[2];
}
