import { GetStaticProps } from "next"

type Category = {
    id: string;
    name: string;
    description: string;
}

export default function CategoryList({ categories }: { categories: Category[] }) {
    return (
        <div>
            <h1>Select Categories</h1>
            <ul>
                {
                    categories.map(category => (
                        // <li key={category.id}><div>{category.name}</div><small>{category.description}</small></li>
                        <li key={category.id}>{category.name}<br />{category.description}</li>
                    ))
                }
            </ul>
        </div>
    )
}

export const getStaticProps: GetStaticProps = async () => {
    const categories: Category[] = [];
    for (let i = 1; i <= 10; i++) {
        categories.push({
            id: `${i}`,
            name: `Category ${i}`,
            description: `Description of category ${i}`
        });
    }

    return {
        props: {
            categories,
        },
        revalidate: 10
    }
}