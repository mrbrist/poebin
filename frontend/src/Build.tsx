import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { getBuild, type buildData } from "./api/getBuild";

async function handleBuild(
    id: string,
    set: React.Dispatch<React.SetStateAction<buildData | undefined>>,
) {
    const build = await getBuild(id);
    if (!build) {
        return;
    }
    set(build);
}

function Build() {
    const navigate = useNavigate();
    const { id } = useParams();

    const [build, setBuild] = useState<buildData | undefined>();

    useEffect(() => {
        if (!id) {
            navigate("/");
            return;
        }

        const init = async () => {
            await handleBuild(id, setBuild);
        };
        init();
    }, []);

    return <div>{build?.Id} page</div>;
}

export default Build;
