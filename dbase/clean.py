import numpy as np
import pandas as pd

import sqlite3 as sql

def clean_escape(text):
    for i in range(len(text) - 1):
        if text[i] == '\r' and text[i + 1].isupper():
            text = text[:i] + '\n' + text[(i + 1):]

    text = text.replace('\r', '')

    text = text.replace(
        '\u2029', '\n\n' # PS
    ).replace(
        '\u2028', '\n' # LS
    )

    return '\n'.join(
        [line.strip() for line in text.split('\n')]
    )

if __name__ == '__main__':
    df = pd.read_csv('PoetryFoundationData.csv')
    df = df[['Title', 'Poet', 'Poem']].copy()

    for (i, r) in df.iterrows():
        df.loc[i, 'Title'] = r['Title'].strip()
        df.loc[i, 'Poem'] = clean_escape(r['Poem'].strip())
        df.loc[i, 'Poet'] = r['Poet'].strip()

    df.rename(
        columns = {
            'Title' : 'title',
            'Poet' : 'author',
            'Poem' : 'poem',
        }, inplace = True
    )

    df['title_noncase'] = [
        title.lower() for title in df['title']
    ]
    df['author_noncase'] = [
        author.lower() for author in df['author']
    ]

    with sql.connect('poetry-database.sqlite3') as conn:
        df.to_sql('Poems', conn)
