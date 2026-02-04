import { useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

function Build() {
    const navigate = useNavigate();
    const { id } = useParams();

    useEffect(() => {
        if (!id) navigate('/', { replace: true });
    }, []);

    return <div>build page</div>;
}

export default Build;
