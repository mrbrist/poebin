import { useState } from "react";
import type { buildData } from "../api/getBuild";
import { parseItem } from "../helpers/item_parser";

function BuildItems({ build }: { build: buildData | undefined }) {
    const itemList = build?.Data?.Items?.ItemList;
    const sets = build?.Data?.Items?.ItemSets;

    if (!itemList || !sets) return null;

    // simple lookup
    const itemMap: Record<string, any> = {};
    itemList.forEach((e: { ID: string | number; Content: any }) => {
        itemMap[e.ID] = parseItem(e.Content);
    });

    const [selectedSetId, setSelectedSetId] = useState(sets[0].ID);
    const selectedSet = sets.find((s: { ID: any }) => s.ID === selectedSetId);

    if (!selectedSet) return null;

    return (
        <div className="block text-left">
            {/* Dropdown */}
            <select
                value={selectedSetId}
                onChange={(e) => setSelectedSetId(Number(e.target.value))}
                className="mb-4 p-2 border rounded"
            >
                {sets.map((set: any) => (
                    <option key={set.ID} value={set.ID}>
                        {set.Title}
                    </option>
                ))}
            </select>

            {/* Items */}
            <div>
                {Object.entries(selectedSet.Gear).map(([slot, id]) => {
                    if (!id || typeof id !== "string") return null;

                    const item = itemMap[id];
                    if (!item) return null;

                    return (
                        <span className="flex max-w-xl items-center">
                            <span className="w-2/8 whitespace-nowrap">
                                {slot}:
                            </span>

                            <span
                                className={`${item.Rarity} flex-1 text-left whitespace-nowrap overflow-hidden text-ellipsis`}
                            >
                                {item.Name}
                            </span>
                        </span>
                    );
                })}
            </div>
        </div>
    );
}

export default BuildItems;
