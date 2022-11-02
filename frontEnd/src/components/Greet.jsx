import { HelloWorldService } from "../services/helloworld-service"
import { useEffect, useState, useCallback } from 'react';

const Greet = () => {
    const [isLoading, setLoading] = useState(true);
    const [greet, setGreet] = useState('');

    const handleGreet = async (e) => {
        let res = await HelloWorldService.getGreet();
        console.log(res);
        setGreet(res.Content);
    }

    useEffect(() => {
        setLoading(true);
        handleGreet().then(() => {
            setLoading(false);
        })
    }, [])

    return (
        <div>
            {!isLoading &&
                <h1>
                    {greet}
                </h1>
            }
        </div>
    )
}

export default Greet