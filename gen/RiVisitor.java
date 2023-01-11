// Generated from java-escape by ANTLR 4.11.1
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link RiParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface RiVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by the {@code Number}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumber(RiParser.NumberContext ctx);
	/**
	 * Visit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMulDiv(RiParser.MulDivContext ctx);
	/**
	 * Visit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAddSub(RiParser.AddSubContext ctx);
	/**
	 * Visit a parse tree produced by {@link RiParser#functionDeclaration}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionDeclaration(RiParser.FunctionDeclarationContext ctx);
	/**
	 * Visit a parse tree produced by {@link RiParser#blockStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBlockStatement(RiParser.BlockStatementContext ctx);
	/**
	 * Visit a parse tree produced by the {@code BlockStms}
	 * labeled alternative in {@link RiParser#blockStatements}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBlockStms(RiParser.BlockStmsContext ctx);
	/**
	 * Visit a parse tree produced by {@link RiParser#start}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStart(RiParser.StartContext ctx);
}