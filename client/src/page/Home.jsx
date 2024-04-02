import { useEffect, useState } from 'react';
import { Link } from "react-router-dom";
import { Col, Container, Row } from 'react-bootstrap';
import Spinner from 'react-bootstrap/Spinner';
import axios from 'axios';

import './../App.css';


function Home() {

    const [apiData, setApiData] = useState(false);
    const [loading, setLoading] = useState(true);

    useEffect(() => {

        const fetchData = async () => {

            try {
                const instance = import.meta.env.VITE_API_ROOT;
                const response = await axios.get(instance);

                if (response.status === 200) {

                    if (response?.data.statusText === "Ok") {
                        setApiData(response?.data.blog_records)
                    }
                }
                setLoading(false);

            } catch (error) {
                setLoading(false);
                console.log("Errorrrrrrr");
            }
        }

        fetchData();
        return () => { };
    }, []);

    if (loading) {
        return (
            <>
                <Container className='spinner'>
                    <Spinner animation="border" role="status">
                        <span className="visually-hidden">Loading...</span>
                    </Spinner>
                </Container>
            </>
        )
    }

    return (
        <Container className="py-4">
            <Row>
                <h3>
                    <Link to="Add" className="btn btn-primary">
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