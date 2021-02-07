import { GetStaticPaths, GetStaticProps } from "next"
import styled from "styled-components"

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




export default function QuizBoard({ quiz }: { quiz: Quiz }) {
    return (
        <div>
            <div>
                Quiz: {quiz.question}
            </div>
            <div>
                Sound:
                <audio controls autoPlay>
                    <source src={quiz.audio.uri} type="audio/ogg" />
                </audio>
                <small>{quiz.audio.id}</small>
            </div>
            <ImageGrid>
                {
                    quiz.images.map(img => (
                        <ImageBox key={img.id}>
                            <Image src={img.uri} alt="description of the image" />
                        </ImageBox>
                    ))
                }
            </ImageGrid>
        </div>
    )
}

export const getStaticProps: GetStaticProps = async ({ params }) => {
    const quizID = params.id;

    const response = await fetch("http://localhost:8080/game/44e420ad-c5fd-4d26-9ed5-2838f2450147");
    const quiz = await response.json();
    // const quiz: Quiz = {
    //     id: "44e420ad-c5fd-4d26-9ed5-2838f2450147",
    //     images: [
    //         {
    //             id: "44e420ad-c5fd-4d26-9ed5-2838f2450147",
    //             uri: "http://localhost:8080/quizzes/44e420ad-c5fd-4d26-9ed5-2838f2450147/image"
    //         },
    //         {
    //             id: "93ac3e8e-e82b-4a4f-adf8-f09d4857b0e0",
    //             uri: "http://localhost:8080/quizzes/93ac3e8e-e82b-4a4f-adf8-f09d4857b0e0/image"
    //         },
    //         {
    //             id: "32ed9040-1c14-4875-9ab3-c504dc7a962c",
    //             uri: "http://localhost:8080/quizzes/32ed9040-1c14-4875-9ab3-c504dc7a962c/image"
    //         },
    //         {
    //             id: "5509e717-a4de-4457-b3b9-5b1676e591e4",
    //             uri: "http://localhost:8080/quizzes/5509e717-a4de-4457-b3b9-5b1676e591e4/image"
    //         }
    //     ],
    //     audio: {
    //         id: "44e420ad-c5fd-4d26-9ed5-2838f2450147",
    //         uri: "http://localhost:8080/quizzes/44e420ad-c5fd-4d26-9ed5-2838f2450147/audio"
    //     }
    // }

    return {
        props: { quiz },
        revalidate: 60,
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
        fallback: false,
    }
}

const ImageGrid = styled.div`
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
    max-width: 1024px;
    max-height: 748px;
    margin: 20px;
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