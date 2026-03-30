import { type buildData } from "../api/getBuild";
import { parseItem } from "../helpers/item_parser";

function BuildMain({ build }: { build: buildData | undefined }) {
    return (
        <div className="col-span-9 rounded-2xl border border-slate-400 p-4">
            {JSON.stringify(parseItem(build?.Data.Items.ItemList[15].Content))}
        </div>
    );
}

export default BuildMain;
