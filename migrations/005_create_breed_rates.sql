-- +goose Up
CREATE TABLE breed_rates (
    id          SERIAL PRIMARY KEY,
    species     VARCHAR(10) NOT NULL,
    breed       VARCHAR(100) NOT NULL,
    size_class  VARCHAR(10) NOT NULL,
    risk_factor DECIMAL(4,2) NOT NULL DEFAULT 1.00,
    UNIQUE(species, breed)
);

-- Popular dog breeds in Japan
INSERT INTO breed_rates (species, breed, size_class, risk_factor) VALUES
('dog', 'トイプードル',             'small',  1.00),
('dog', 'チワワ',                   'small',  1.05),
('dog', 'ミニチュアダックスフンド',   'small',  1.15),  -- prone to back issues
('dog', 'ポメラニアン',             'small',  1.00),
('dog', 'ヨークシャーテリア',       'small',  1.00),
('dog', 'シーズー',                 'small',  1.05),
('dog', 'マルチーズ',               'small',  1.00),
('dog', 'ミニチュアシュナウザー',    'small',  1.05),
('dog', 'フレンチブルドッグ',       'medium', 1.40),  -- high medical costs
('dog', '柴犬',                     'medium', 1.10),
('dog', 'コーギー',                 'medium', 1.15),
('dog', 'ビーグル',                 'medium', 1.05),
('dog', 'ボーダーコリー',           'medium', 1.05),
('dog', 'ゴールデンレトリバー',     'large',  1.20),
('dog', 'ラブラドールレトリバー',    'large',  1.15),
('dog', 'シベリアンハスキー',       'large',  1.10),
('dog', 'バーニーズマウンテンドッグ', 'large', 1.35),  -- shorter lifespan, high risk
('dog', 'ミックス犬',               'small',  0.90),  -- mixed breeds healthier
-- Cats
('cat', 'スコティッシュフォールド',  'small',  1.20),  -- prone to joint issues
('cat', 'マンチカン',               'small',  1.15),
('cat', 'アメリカンショートヘア',    'small',  1.00),
('cat', 'ブリティッシュショートヘア', 'small', 1.05),
('cat', 'ラグドール',               'small',  1.05),
('cat', 'ノルウェージャンフォレストキャット', 'small', 1.10),
('cat', 'ペルシャ',                 'small',  1.15),
('cat', 'メインクーン',             'small',  1.10),
('cat', 'ミックス猫',               'small',  0.85);  -- mixed breeds healthier

-- +goose Down
DROP TABLE breed_rates;