import { GetStaticPaths, GetStaticProps } from "next"
import { NextRouter, useRouter } from "next/router"
import { Dispatch, SetStateAction, useState } from "react"
import styled from "styled-components"
import MessageModal from "../../../components/MessageModal"

type MediaInfo = {
    id: string;
    uri: string;
}

type Quiz = {
    question: string;
    images: MediaInfo[];
    audio: MediaInfo;

}

//TODO Add Player and fetch it.
type Game = {
    id: string;
    quizzes: string[];
}

const validateSelection = () => {
    return true;
}

const onSelectImage = (target: string, updateFn: Dispatch<SetStateAction<boolean>>) => {
    console.log("Clicked:", target);

    updateFn(target == "44e420ad-c5fd-4d26-9ed5-2838f2450147");
}

const onCloseSuccessModal = (router: NextRouter, setSuccess: Dispatch<SetStateAction<boolean>>) => {
    setSuccess(false);

    const next_quiz = "b";

    router.push(`/game/quizzes/${next_quiz}`);
}

export default function QuizBoard({ quiz, nextQuiz }: { quiz: Quiz, nextQuiz: string }) {
    const [success, setSuccess] = useState(false);
    const router = useRouter();

    const next_quiz = "b";

    return (
        <Container>
            <div>
                <h1>{quiz.question}</h1>
                Result: {success ? "Yes!" : "Ouch!"}
            </div>
            <AudioBox>
                <audio controls autoPlay>
                    <source src={quiz.audio.uri} type="audio/ogg" />
                </audio>
            </AudioBox>
            <ImageGrid>
                {
                    quiz.images.map(img => (
                        <ImageBox
                            key={img.id}
                            style={{ backgroundColor: success && img.id == quiz.audio.id ? "green" : "" }}>
                            <Image
                                src={img.uri}
                                alt="description of the image"
                                onClick={(event) => onSelectImage(img.id, setSuccess)} />
                        </ImageBox>
                    ))
                }
            </ImageGrid>
            <MessageModal
                message="Parabéns, você acertou! :)"
                visible={success}
                onClose={() => {setSuccess(false);router.push(nextQuiz);}} />
        </Container>
    )
}

CREATE A STATE ENGINE OF THE GAMEPLAY TO MANAGE THE QUIZZES THAT SHOULD BE PRESENTED.
THE GAME MUST RETURN THE FIRST QUIZ WHEN IT STARTS.
THUS, IT SHOULD HAVE A VALIDATION ENDPOINT THAT EVALUATES THE USER RESPONSE OF THE QUIZ AND THEN 
RETURN THE NEXT QUIZ ID UNTIL THE GAME FINISHES.

export const getStaticProps: GetStaticProps = async ({ params }) => {
    const quizID = params.id;
    console.log("quizID:", quizID);

    // const response = await fetch("http://localhost:8080/game/44e420ad-c5fd-4d26-9ed5-2838f2450147");
    // const quiz = await response.json();

    const quiz = quizMock;

    const nextQuizMap = {
        "a": "/game/quizzes/b",
        "b": "/game/quizzes/c",
        "c": "/game/quizzes/a"
    };

    return {
        props: {
            quiz: quiz,
            nextQuiz: nextQuizMap[quizID as string]
        },
        // revalidate: 1,
    }

}

export const getStaticPaths: GetStaticPaths = async () => {

    const game: Game = {
        id: "1000",
        quizzes: ["a", "b", "c"],
    };
    // const paths = game.quizzes.map(quiz => {
    //     params: { 
    //         id: `/game/gameplay/quizzes/${quiz}` 
    //     }
    // });

    const paths = game.quizzes.map(quiz => `/game/quizzes/${quiz}`);

    return {
        paths,
        fallback: 'blocking',
    }
}

const Container = styled.div`
    display: flex;
    /* flex-wrap: wrap; */
    flex-direction: column;
    justify-content: space-around;
    max-width: 1024px;
    max-height: 748px;
    margin: 20px;
`

const ImageGrid = styled.div`
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
    border: 1px solid #7e7575;
    border-radius: 16px;
`

const ImageBox = styled.div`
    :hover {
        cursor: pointer;
        -webkit-box-shadow: 0px 0px 14px 2px rgba(0,0,0,0.5);
        -moz-box-shadow: 0px 0px 14px 2px rgba(0,0,0,0.5);
        box-shadow: 0px 0px 14px 2px rgba(0,0,0,0.5);
    };
    
    width: 220px;
    height: 220px;
    margin: 10px;
    padding: 4px;
    border-radius: 4px;
`

const Image = styled.img`
    max-width: 200px;
    max-height: 200px;
    
`

const AudioBox = styled.div`
    margin: 10px;
    padding: 4px;
    border-radius: 4px;
    border: 1px solid #7e7575;
`

const quizMock: Quiz = {
    "question": "What's this animal?",
    "images": [
        {
            id: "44e420ad-c5fd-4d26-9ed5-2838f2450147",
            uri: "http://localhost:8080/quizzes/44e420ad-c5fd-4d26-9ed5-2838f2450147/image"
        },
        {
            id: "93ac3e8e-e82b-4a4f-adf8-f09d4857b0e0",
            uri: "http://localhost:8080/quizzes/93ac3e8e-e82b-4a4f-adf8-f09d4857b0e0/image"
        },
        {
            id: "32ed9040-1c14-4875-9ab3-c504dc7a962c",
            uri: "http://localhost:8080/quizzes/32ed9040-1c14-4875-9ab3-c504dc7a962c/image"
        },
        {
            id: "5509e717-a4de-4457-b3b9-5b1676e591e4",
            uri: "http://localhost:8080/quizzes/5509e717-a4de-4457-b3b9-5b1676e591e4/image"
        }
    ],
    audio: {
        id: "44e420ad-c5fd-4d26-9ed5-2838f2450147",
        uri: "http://localhost:8080/quizzes/44e420ad-c5fd-4d26-9ed5-2838f2450147/audio"
    }
}