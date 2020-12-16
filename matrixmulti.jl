function multi(A::Matrix, B ::Matrix)
    if size(A, 2) != size(B, 1)
        throw(ArgumentError("matrix A doesn't match matrix B"))
    end
    temp = zeros(size(A, 1), size(B, 2))
    for i = 1:size(A, 1)
        for j = 1:size(B, 2)
            temp[i, j] = sum(A[i, :] .* B[:, j])
        end
    end
    return temp
end

ta = [1 2 3 ; 4 5 6];
tb = [1 2 ; 3 4 ; 5 6];
multi(ta, tb)