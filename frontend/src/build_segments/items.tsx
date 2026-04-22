import type { buildData } from "../api/getBuild";
import { parseItem } from "../helpers/item_parser";

function BuildItems({ build }: { build: buildData | undefined }) {
    return (
        <div>
            {build?.Data.Items.ItemList.map(
                (item: { Content: any }, index: number) => (
                    <span key={index} className="block">
                        {parseItem(item.Content).Name}
                    </span>
                ),
            )}
        </div>
    );
}

export default BuildItems;
