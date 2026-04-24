export interface itemData {
    Name: string;
    Base: string;
    Rarity: string;
    Quality: number;
    Sockets: string;
    ItemLevel: number;
    Mods: [string];
}

export function parseItem(raw: any) {
    raw = String(raw);
    raw = raw.replace(/\t/g, "");

    let itemSplit = raw.split("\n");

    let item: itemData = {
        Name: itemSplit[2],
        Base: itemSplit[3],
        Rarity: itemSplit[1].split(" ")[1].toLowerCase(),
        Quality: itemSplit[10],
        Sockets: itemSplit[11],
        ItemLevel: itemSplit[9],
        Mods: itemSplit.slice(14, itemSplit.length).filter((n: any) => n),
    };
    return item;
}
