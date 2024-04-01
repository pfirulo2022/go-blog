import axios from "axios";
import { useState, useEffect} from "react";
import { Container,Row, Col } from "react-bootstrap";
import { useParams } from "react-router-dom";

 

function Blog() {
    const params = useParams();
    console.log(params.id);
    const [apiData, setApiData] = useState(false);

    useEffect(() => {

        const fetchData = async () => {

            try {

                const instance = import.meta.env.VITE_API_ROOT +"/"+ params.id;
                console.log(instance);
                const response = await axios.get(instance)
                if (response.status === 200) {

                    if (response?.data.statusText === "Ok") {
                        setApiData(response?.data?.record)
                    }
                }

            } catch (error) {
                console.log("Error detail");
            }
        }

        fetchData();
        return () => { };
 
    }, [params]);

    return (
        <Container>
            <Row>                
                <Col xs="12"><h1>{apiData.title}</h1></Col>

                <Col xs="12" ><p>{apiData.post}</p></Col>
            </Row>
             </Container>
    )
}
export default Blog;