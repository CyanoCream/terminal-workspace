-- Sample terminals
INSERT INTO terminals (name, address, latitude, longitude) VALUES
                                                               ('Terminal A', 'Jl. Terminal A No.1', -6.175392, 106.827153),
                                                               ('Terminal B', 'Jl. Terminal B No.2', -6.186486, 106.834091),
                                                               ('Terminal C', 'Jl. Terminal C No.3', -6.221762, 106.845733),
                                                               ('Terminal D', 'Jl. Terminal D No.4', -6.202393, 106.832720),
                                                               ('Terminal E', 'Jl. Terminal E No.5', -6.189957, 106.821857);

-- Sample gates for Terminal A
INSERT INTO gates (terminal_id, name, is_active) VALUES
                                                     (1, 'Gate A1', true),
                                                     (1, 'Gate A2', true),
                                                     (1, 'Gate A3', true);

-- Sample pricing
INSERT INTO terminal_distances (from_terminal_id, to_terminal_id, distance, base_price) VALUES
                                                                                            (1, 2, 5.2, 10000),
                                                                                            (1, 3, 8.7, 15000),
                                                                                            (1, 4, 6.3, 12000),
                                                                                            (1, 5, 4.8, 9000),
-- Tambahkan kombinasi lainnya
                                                                                            (2, 3, 7.1, 13000),
                                                                                            (2, 4, 5.5, 11000);