from solver import Solver

solver = Solver(num_ant=30, num_node=300, q=100, alpha=3, beta=5, rou=0.9,
                max_iteration=300, min_tau=0, max_tau=3000,
                border_unchanged=30, start_node=0,
                path="./data/qiita_data.json")

solver.run_aco()
