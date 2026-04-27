import type { buildData } from "../api/getBuild";

function BuildTree({ build }: { build: buildData | undefined }) {
    return (
        <div>
            {build?.Data.Tree.Specs[0].Title}
            <br />
            {build?.Data.Tree.Specs[0].Nodes}
        </div>
    );
}

export default BuildTree;
