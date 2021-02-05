import { GetStaticPaths, GetStaticProps } from "next"

type Quiz = {
    id: string;
    images: string[];
    audio: string;

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
                Quiz: {quiz.id}
            </div>
            <div>
                Sound: <audio src={quiz.audio}></audio>
                <small>{quiz.audio}</small>
            </div>
            <div>
                {
                    quiz.images.map(img => (
                        <div key={img}>
                            <div>
                                <img src={img} alt="description of the image" />
                            </div>
                            <small>{img}</small>
                        </div>
                    ))
                }
            </div>
        </div>
    )
}

export const getStaticProps: GetStaticProps = async ({ params }) => {
    const quizID = params.id;

    const quiz: Quiz = {
        id: "a",
        audio: "a/audio",
        images: [
            "a/images",
            "b/images",
            "c/images",
        ]
    }

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