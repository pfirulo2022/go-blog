import { useEffect, useState } from 'react';
import { Link } from "react-router-dom";
import { Col, Container, Row } from 'react-bootstrap';
import axios from 'axios';

import './../App.css';


function Home() {

    const [apiData, setApiData] = useState(false);

    useEffect(() => {

        const fetchData = async () => {

            try {

                const instance = import.meta.env.VITE_API_ROOT;


                const response = await axios.get(instance)
                if (response.status === 200) {

                    if (response?.data.statusText === "Ok") {
                        setApiData(response?.data.blog_records)
                    }
                }

            } catch (error) {
                console.log("Errorrrrrrr");
            }
        }

        fetchData();
        return () => { };
    }, []);

    return (
        <Container className="py-2">
            <Row>
                <h3>
                    <Link to="add" className="btn btn-primary">
                        Add New
                    </Link>
                </h3>

                <h5>{location.state && location.state}</h5>

                {apiData &&
                    apiData.map((record, index) => (
                        <Col key={index} xs="3" className="py-2 box">
                            <div className="img-box justify-content-center py-2 mb-3">
                                <img width="150" height="150" src={`${import.meta.env.VITE_API_ROOT}/${record.image}`} />
                            </div>
                            <div className="title">
                                <Link to={`blog/${record.id}`}> {record.title}</Link>
                            </div>

                            <div>
                                <Link to={`edit/${record.id}`}>
                                    <i className="fa fa-solid fa-pencil fa-1x" />
                                </Link>
                                &nbsp;
                                <Link to={`delete/${record.id}`}>
                                    <i className="fa fa-solid fa-trash fa-1x" />
                                </Link>
                            </div>
                            <div>{record.post}</div>
                        </Col>
                    ))}
            </Row>
        </Container>
    );
}
export default Home;