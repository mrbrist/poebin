import type { buildData } from "../api/getBuild";
import { parseItem } from "../helpers/item_parser";

function BuildItems({ build }: { build: buildData | undefined }) {
    const itemList = build?.Data?.Items?.ItemList;
    const sets = build?.Data?.Items?.ItemSets;

    if (!itemList || !sets) return null;

    const items = itemList.map((e: { Content: any }) => parseItem(e.Content));

    return (
        <div className="block text-left">
            {items.map((item: any, index: any) => (
                <span key={item.ID ?? index} className={`block ${item.Rarity}`}>
                    {item.Name}
                </span>
            ))}
        </div>
    );
}

export default BuildItems;
