

interface IItem{
    from : number,
    to : number,
    amount : number
}

export default function Item({from,to,amount}: IItem){
    return(
        <h3>
            {from}{'---->'}{to} {'$'}{amount}
        </h3>
    )
}